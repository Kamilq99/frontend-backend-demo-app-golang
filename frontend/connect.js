function GetText() {
    fetch("http://backend:8080/text") // zamiast localhost
        .then(response => response.json())
        .then(data => {
            document.getElementById("textarea").value = data.text;
        })
        .catch(error => console.error(error));
}
