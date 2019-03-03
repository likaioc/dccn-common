import nacl from "tweetnacl";
import get_balance from "./get_balance";
import { sha256 } from 'js-sha256';

let bc_addr = '//localhost:26657/abci_query';
export const set_blockchain_addr = (blockchain_addr) => {
    bc_addr = '//' + blockchain_addr + '/abci_query';
}

function hex2bytes(hex)
{
    var bytes = [], str;

    for(var i=0; i< hex.length-1; i+=2)
        bytes.push(parseInt(hex.substr(i, 2), 16));

    return bytes
}

export default async (from, to, amount, private_key, public_key) => {

    const balance = await get_balance(from); //TODO check balance is FAIL.

    // handle nonce
    let nonce = balance.split(":")[1];
    let nonceInt = parseInt(nonce); //TODO check parseInt fail.
    nonceInt++

    // for test
    //nonceInt = 10 

    // handle hash
    const sha245_sum = sha256(from + to + amount + nonceInt);
    let arr = new Uint8Array(32);
    arr.set(hex2bytes(sha245_sum));

    // handle priv key
    var raw = atob(private_key); //TODO check DOMException
    var rawLength = raw.length;
    var pk_array = new Uint8Array(new ArrayBuffer(rawLength));
    var i = 0;
    for(i = 0; i < rawLength; i++) {
        pk_array[i] = raw.charCodeAt(i);
    }

    // handle sign
    let sign_array = nacl.sign(arr, pk_array)
    const sign = btoa(String.fromCharCode.apply(null, sign_array.slice(0,64)));

    await fetch(bc_addr + '?tx="trx_send=' + from + ":" + to + ":" + amount + ":" + nonceInt + ":" + public_key + ":" + sign + '"', {
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
