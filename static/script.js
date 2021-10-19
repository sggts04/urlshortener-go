let form = document.getElementById('form');

form.addEventListener('submit', function(e) {
    e.preventDefault();
    let longURL = document.getElementById('longURL').value;
    fetch('/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: `longURL=${longURL}`
        })
        .then(response => response.json())
        .then(data => {
            let id = data.id;
            let url = window.location.href;
            let shortURL = url.substring(0, url.lastIndexOf('/') + 1) + id;
            console.log(shortURL);
        });
});