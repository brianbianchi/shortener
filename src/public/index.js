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
            let text =
                key === 'created' || key === 'last_visited'
                    ? document.createTextNode(new Date(element[key]).toLocaleString())
                    : document.createTextNode(element[key]);
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
        .then(response => response.json())
        .then(json => {
            if (json.Status === 400 || json.Status === 500) {
                document.getElementById('error').innerHTML = json.Message;
                return;
            }
            document.getElementById('link').value = "";
            document.getElementById('error').innerHTML = "";
            const url = apiBaseUrl + '/' + json.Url.code;
            document.getElementById('msg').innerHTML =
                `<a href='${url}'>${json.Url.link}</a> <button onclick="toClipboard('${url}')">Copy shortened URL</button>`;
            if (json.Status === 201) {
                let table = document.querySelector("table");
                let row = table.insertRow(1);
                for (key in json.Url) {
                    let cell = row.insertCell();
                    let text =
                        key === 'created' || key === 'last_visited'
                            ? document.createTextNode(new Date(json.Url[key]).toLocaleString())
                            : document.createTextNode(json.Url[key]);
                    cell.appendChild(text);
                }
            }
        });
}

function getUrls() {
    fetch(apiBaseUrl + '/urls/')
        .then(response => response.json())
        .then(json => {
            if (json.Status !== 200) {
                document.getElementById('error').innerHTML = json.Message;
                return;
            }
            let table = document.querySelector("table");
            let data = Object.keys(json.Urls[0]);
            generateTableHead(table, data);
            generateTable(table, json.Urls);
        });
}

function toClipboard(str) {
    navigator.clipboard.writeText(str);
}