document.addEventListener('DOMContentLoaded', function () {
    function showResult() {
        const result = document.getElementById('result');
        const error = document.getElementById('error');
        
        result.classList.remove('hidden');
        error.classList.add('hidden');
    }

    function showError(message) {
        const result = document.getElementById('result');
        const error = document.getElementById('error');
        
        result.classList.add('hidden');
        error.innerText = message || "Invalid URL";
        error.classList.remove('hidden');
    }

    document.querySelector('form').addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const url = document.getElementById('url').value;
        
        if (isValidUrl(url)) {
            try {
                const response = await fetch("/short", {
                    method: "POST",
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ url })
                });

                if (response.ok) {
                    const res = await response.json();
                    document.getElementById('shortened-url').value = res.shortURL;
                    showResult();
                } else {
                    showError(response.statusText);
                }
            } catch (error) {
                showError("An error occurred while shortening the URL.");
                console.error('Error:', error);
            }
        } else {
            showError();
        }
    });

    function isValidUrl(string) {
        try {
            new URL(string);
            return true;
        } catch (error) {
            return false;
        }
    }
});

function copyToClipboard() {
    const shortenedUrl = document.getElementById('shortened-url');
    
    if (navigator.clipboard) {
        navigator.clipboard.writeText(shortenedUrl.value)
            .then(() => {
                console.log('Text copied to clipboard');
            })
            .catch(err => {
                console.error('Failed to copy text: ', err);
            });
    } else {
        shortenedUrl.select();
        document.execCommand('copy');
        console.warn('Fallback: execCommand used.');
    }
}
