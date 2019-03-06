let bc_addr = '//localhost:26657/tx_search';
export const set_blockchain_addr = (blockchain_addr) => {
    bc_addr = '//' + blockchain_addr + '/tx_search';
}
const from_addr_base64 = btoa('app.fromaddress');
const to_addr_base64 = btoa('app.toaddress');
const timestamp_base64 = btoa('app.timestamp');

const get = async (uri, wallet_addr, historys) => {
    await fetch(bc_addr + '?query="' + uri + '=\'' + wallet_addr.toString() + '\'"&prove=true', {
        method: 'GET',
        mode: 'cors',
    }).then(res => {
        return res.json();
    }).then(json => {
        if (json['error']) {
            throw json['error'];
        } else if (json.result['txs'] == undefined) {
            console.log(json.result);
            throw "unknown error";
        }
        return json.result.txs;
    }).then(txs => {
        txs.forEach(tx => {
            let history = { amount: atob(tx.tx).split(':')[2] };
            tx.tx_result.tags.forEach(tag => {
                switch (tag.key) {
                    case from_addr_base64:
                        history.from_addr = atob(tag.value);
                        break;
                    case to_addr_base64:
                        history.to_addr = atob(tag.value);
                        break;
                    case timestamp_base64:
                        history.timestamp = atob(tag.value);
                }
            })
            historys.push(history);
        });
    }).catch(err => {
        console.log('请求错误', err);
    });
}

export const get_send_history = async (wallet_addr) => {
    let historys = new Array;
    await get('app.fromaddress', wallet_addr, historys)
    return historys;
}

export const get_recv_history = async (wallet_addr) => {
    let historys = new Array;
    await get('app.toaddress', wallet_addr, historys)
    return historys;
}

export const get_history = async (wallet_addr) => {
    let historys = new Array;
    await get('app.fromaddress', wallet_addr, historys)
    await get('app.toaddress', wallet_addr, historys)
    return historys;
}