console.log("Welcome to the Tour of Pirsch Analytics!");
console.log("Please check out the tour on our website to learn more: https://pirsch.io/tour");

// Before attaching any Pirsch events, we wait for the page to be fully loaded.
document.addEventListener("DOMContentLoaded", () => {
    // The data-use-backend attribute will be present on the script if we track from the server-side.
    // If that's the case, we can skip tracking form submissions using JavaScript and use a different event function.
    // But we need to track outbound links now, as that isn't picked up by the Pirsch snippet.
    const useBackend = !!document.querySelector("script[data-use-backend]");

    if (!useBackend) {
        trackFormSubmissions();
    } else {
        trackOutboundLinkClicks();
    }

    trackScrollDepth(useBackend);
});

function trackFormSubmissions() {
    // Find all forms with the data-pirsch-form attribute and use the value as the event name.
    document.querySelectorAll("[data-pirsch-form]").forEach(form => {
        let preventSubmission = true;

        form.addEventListener("submit", e => {
            // Prevent submitting the form before we've sent the event.
            if (preventSubmission) {
                e.preventDefault();
                preventSubmission = false;
            }

            // Extract all input fields with the data-pirsch-input.
            const meta = {};
            form.querySelectorAll("[data-pirsch-input]")
                .forEach(i => meta[i.getAttribute("name")] = i.value);

            // Send it to Pirsch and re-trigger the event to submit the form.
            pirsch(form.getAttribute("data-pirsch-form"), {meta}).finally(() => e.target.submit());
        });
    });
}

function trackOutboundLinkClicks() {
    // Find all links on the page.
    const links = document.getElementsByTagName("a");

    // Filter links with the data-pirsch-ignore attribute or pirsch-ignore class name.
    for (const link of links) {
        if (!link.hasAttribute("data-pirsch-ignore") && !link.classList.contains("pirsch-ignore")) {
            // Filter internal links.
            const url = new URL(link.href);

            if (url !== null && url.hostname !== location.hostname) {
                // Add event listeners for click and middle mouse button.
                link.addEventListener("click", () => trackEvent("Outbound Link Click", {url: url.href}));
                link.addEventListener("auxclick", () => trackEvent("Outbound Link Click", {url: url.href}));
            }
        }
    }
}

function trackScrollDepth(useBackend) {
    // List of pages we would like to track the scroll depth on.
    const pages = [
        "/phone",
        "/pad",
        "/watch"
    ];

    if (pages.includes(location.pathname)) {
        // Update the scroll position.
        let position = 0;

        window.addEventListener("scroll", () => {
            const p = getScrollPercent();

            if (p > position) {
                position = p;
            }
        });

        // Before we leave the page, send it.
        window.onbeforeunload = beforeUnload;

        function beforeUnload() {
            useBackend ? trackEvent("Scroll Depth", {position}) : pirsch("Scroll Depth", {meta: {position}});
        }

        function getScrollPercent() {
            const h = document.documentElement,
                b = document.body,
                st = 'scrollTop',
                sh = 'scrollHeight';
            return Math.floor((h[st] || b[st]) / ((h[sh] || b[sh]) - h.clientHeight) * 100);
        }
    }
}

function trackEvent(name, meta) {
    // We can only track strings...
    for (let key in meta) {
        meta[key] = meta[key].toString();
    }

    // Use sendBeacon instead of fetch, so that the event goes through even if we leave the page.
    navigator.sendBeacon("/p/event", JSON.stringify({
        name,
        meta,
        path: location.pathname
    }));
}
