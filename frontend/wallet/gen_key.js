import session25519 from 'session25519';
import sha256 from "./sha256";

const sum = (email, password, callback) => {
    let keys = {};

    session25519(email, password, (err, keys) => {
        if (err !== null) {
            throw err;
        }

        callback(keys)
    });
}

const gen_key = (email, password, callback) => {
    sum(email, password, (keys) => {
        const addr = sha256(keys.publicSignKey).substr(0, 20);

        callback({
            public_key: keys.publicSignKeyBase64,
            private_key: keys.secretSignKeyBase64,
            addr: addr,
        })
    });
};

export default gen_key;