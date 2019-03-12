
export const get_nonce = async (bc_addr, wallet_addr) => {
    let nonce = 'FAIL';
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
        nonce = atob(val).split(':', 2)[1];
    }).catch(err => {
        console.log('request error:', err);
    });

    return nonce;
}
