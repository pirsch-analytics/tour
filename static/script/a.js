console.log("Welcome to the Tour of Pirsch Analytics!");
console.log("Please check out the tour on our website to learn more: https://pirsch.io/tour");

// Before attaching any Pirsch events, we wait for the page to be fully loaded.
document.addEventListener("DOMContentLoaded", () => {
    trackFormSubmissions();
    trackScrollDepth();
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

function trackScrollDepth() {
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
            pirsch("Scroll Depth", {meta: {position}});
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
