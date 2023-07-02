const apiUrl = 'https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1'
let currencies;

let res = fetch(apiUrl).then(response => {
    if(!response.ok){
        throw new Error('Network response failed');
    }
    return response.json()
}).then(data => {
    currencies = data;
    updateCurrencies();
}).catch(error => {
    console.error('Error:', error);
})


function updateCurrencies() {
    let table = document.createElement('table');
    let headerRow = initHeaders();
    table.appendChild(headerRow);
    for(let i = 0; i < currencies.length; i++){
        let row = document.createElement('tr');
        let cellId = document.createElement('td');
        cellId.textContent = currencies[i]['id'];
        row.appendChild(cellId);

        let cellSymbol = document.createElement('td');
        cellSymbol.textContent = currencies[i]['symbol'];
        row.appendChild(cellSymbol);

        let cellName = document.createElement('td');
        cellName.textContent = currencies[i]['name'];
        row.appendChild(cellName);
        if(i < 5){
            row.style.backgroundColor = 'skyblue';
        }
        if(currencies[i]['symbol'] === 'usdt'){
            row.style.backgroundColor = 'green';
        }
        table.appendChild(row);

    }

    document.getElementById("currencies").appendChild(table);
}


function initHeaders(){
    let headerRow = document.createElement('tr');
    let headerCellId  = document.createElement('th')
    headerCellId.textContent = 'id';
    headerRow.appendChild(headerCellId);

    let headerCellSymbol  = document.createElement('th')
    headerCellSymbol.textContent = 'symbol';
    headerRow.appendChild(headerCellSymbol);

    let headerCellName  = document.createElement('th')
    headerCellName.textContent = 'name';
    headerRow.appendChild(headerCellName);
    return headerRow;
}
