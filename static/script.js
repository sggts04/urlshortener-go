let form = document.getElementById('form');
let clicker = document.getElementById('expand-click');
let parent = document.getElementById('form-res');
let customIdInput = document.getElementById('customId');
let longURLInput = document.getElementById('longURL');

isValidURL = str => {
    var a  = document.createElement('a');
    a.href = str;
    return (a.host && a.host != window.location.host);
}

clicker.addEventListener('click', () => {
    clicker.style.display = 'none';
    customIdInput.style.display = 'block';
});

form.addEventListener('submit', (e) => {
    e.preventDefault();
    let longURL = longURLInput.value;
    let customId = customIdInput.value;
    if(!isValidURL(longURL)) {
        alert('Invalid URL');
        return;
    }
    let body = `longURL=${longURL}`
    if(customId.trim().length > 0)
        body += `&customId=${customId}`
    fetch('/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: body
        })
        .then(async (response) => {
            if(response.status == 201)
                return response.json()
            else {
                let res = await response.json();
                alert(res.message);
                return;
            }
        })
        .then(data => {
            if(!data) return;
            let id = data.id;
            let url = window.location.href;
            let shortURL = url.substring(0, url.lastIndexOf('/') + 1) + id;
            parent.innerHTML = `<p class="result">Your short URL is: <a href="${shortURL}">${shortURL}</a></p>`;
        });
});