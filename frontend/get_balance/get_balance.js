let bc_addr = '//localhost:26657/abci_query';
const set_blockchain_addr = (blockchain_addr) => {
    bc_addr = '//' + blockchain_addr + '/abci_query';
}

const get_balance = async (wallet_addr) => {
    let balance = 'FAIL';
    await fetch(bc_addr + '?data="bal:' + wallet_addr.toString() + '"', {
        method: 'GET',
        mode: 'cors',
    }).then(res => {
        return res.json();
    }).then(json => {
        if (json['error']) {
            throw json['error'];
        }
        return json.result.response;
    }).then(data => {
        return data.value;
    }).then(val => {
        balance = atob(val);
    }).catch(err => {
        console.log('请求错误', err);
    });

    return balance;
}

// export default {
//     set_blockchain_addr: set_blockchain_addr,
//     get_balance: get_balance,
// };