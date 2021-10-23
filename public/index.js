const apiBaseUrl = 'http://127.0.0.1:8000';

const urlParams = new URLSearchParams(window.location.search);
const myParam = urlParams.get('q');

if (myParam) {
    redirect(myParam);
}
getShortens();

function redirect(code) {
    const url = apiBaseUrl + '/api/urls/' + code;

    fetch(url, { method: 'GET' })
        .then(response => response.json())
        .then(json => {
            if (json.link) {
                window.location.href = json.link;
            }
        });
}

function generateTableHead(table, data) {
    let thead = table.createTHead();
    let row = thead.insertRow();
    for (let key of data) {
        let th = document.createElement("th");
        let text = document.createTextNode(key);
        th.appendChild(text);
        row.appendChild(th);
    }
}

function generateTable(table, data) {
    for (let element of data) {
        let row = table.insertRow();
        for (key in element) {
            let cell = row.insertCell();
            let text = document.createTextNode(element[key]);
            cell.appendChild(text);
        }
    }
}

function sendLink() {
    const link = document.getElementById('link').value;
    const bodyObj = {
        "link": link
    };

    fetch(apiBaseUrl + '/api/urls', {
        method: 'POST',
        mode: 'no-cors',
        body: JSON.stringify(bodyObj),
        headers: {
            'Content-type': 'application/json;'
        }
    })
        .then(response => {
            console.log(response);
            response.json();
        })
        .then(json => {
            console.log(json);
        });
}

function getShortens() {
    fetch(apiBaseUrl + '/api/urls')
        .then(response => response.json())
        .then(json => {
            let table = document.querySelector("table");
            let data = Object.keys(json[0]);
            generateTableHead(table, data);
            generateTable(table, json);
        });
}
