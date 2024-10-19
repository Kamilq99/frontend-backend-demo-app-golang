function GetText() {
    fetch("/text")
        .then(response => response.json())
        .then(data => {
            document.getElementById("textarea").value = data.text
        })
        .catch(error => console.error(error));
}