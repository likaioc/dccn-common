import session25519 from 'session25519';
import { sha256 } from 'js-sha256';


const sum = (email, password, callback) => {
    let keys = {};

    session25519(email, password, (err, keys) => {
        if (err !== null) {
            throw err;
        }

        callback(keys)
    });
}

const toHex = (str) => {
    var result = '';
    for (var i = 0; i < str.length; i++) {
        result += str.charCodeAt(i).toString(16);
    }
    return result;
}

const gen_key = (email, password, callback) => {
    sum(email, password, (keys) => {
        console.log(sha256(keys.publicSignKey).substr(0, 40))
        const addr = sha256(keys.publicSignKey).substr(0, 40).toUpperCase();

        callback({
            public_key: keys.publicSignKeyBase64,
            private_key: keys.secretSignKeyBase64,
            addr: addr,
        })
    });
};

export default gen_key;
