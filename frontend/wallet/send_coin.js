import nacl from "tweetnacl";
import get_balance from "./get_balance";
import { sha256 } from 'js-sha256';

let bc_addr = '//localhost:26657/abci_query';
export const set_blockchain_addr = (blockchain_addr) => {
    bc_addr = '//' + blockchain_addr + '/abci_query';
}

export default async (from, to, amount, private_key, public_key) => {

    const balance = await get_balance(from);
    console.log(balance)
    let nonce = balance.split(":")[1];
    nonce++;

    const sha245_sum = sha256(from + to + amount + nonce);
    console.log(1)

    let arr = new Uint8Array(32);
    arr.set(sha245_sum.substr(0, 32));
    console.log(2)

    let key_arr = new Uint8Array(64);
    console.log(atob(private_key).length)

    key_arr.set(atob(private_key));
    console.log(3)
    const sign = nacl.sign(arr, key_arr);
    console.log(4)

    await fetch(bc_addr + '?tx="trx_send=' + from + ":" + to + ":" + amount + ":" + nonce + ":" + public_key + ":" + sign + '"', {
        method: 'GET',
        mode: 'cors',
    }).then(res => {
        console.log(res)
        return res.json();
    }).then(json => {
        if (json['error']) {
            throw json['error'];
        }
        return json.result.response;
    }).then(data => {
        return data.value;
    }).then(val => {
        balance = btoa(val);
    }).catch(err => {
        console.log('请求错误', err);
    });
}