const apiBaseUrl = 'http://localhost:8000';
getUrls();

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

async function postUrl() {
    const link = document.getElementById('link').value;
    const bodyObj = {
        "link": link
    };

    await fetch(apiBaseUrl + '/urls/', {
        headers: {
            'Content-Type': 'application/json'
        },
        method: 'POST',
        body: JSON.stringify(bodyObj),
    })
        .then(response => {
            response.json();
        })
        .then(json => {
            console.log(json);
        });
}

function getUrls() {
    fetch(apiBaseUrl + '/urls/')
        .then(response => response.json())
        .then(json => {
            let table = document.querySelector("table");
            let data = Object.keys(json[0]);
            generateTableHead(table, data);
            generateTable(table, json);
        });
}
