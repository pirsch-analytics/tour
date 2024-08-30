console.log("Welcome to the Tour of Pirsch Analytics!");
console.log("Please check out the tour on our website to learn more: https://pirsch.io/tour");

// Before attaching any Pirsch events, we wait for the page to be fully loaded.
document.addEventListener("DOMContentLoaded", () => {
    trackFormSubmissions();
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
