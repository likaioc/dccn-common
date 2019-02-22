import gen_key from "./gen_key";
import get_balance from "./get_balance";
import send_coin from "./send_coin";

let element = null;

const echo = (text) => {
    element.appendChild(document.createElement("br"))
    if (text != null) {
        element.appendChild(document.createTextNode(text))
    }
}

document.getElementById("gen_keys_btn").onclick = () => {
    gen_key('me@example.com', 'brig alert rope welsh foss rang orb', (key) => {
        element = document.getElementById("gen_keys");
        echo();
        echo("public key:");
        echo(key.public_key);
        echo();
        echo("secret key:");
        echo(key.private_key);
        echo();
        echo("wallet address:");
        echo(key.addr);
        echo();
    });
}

document.getElementById("get_balance_btn").onclick = async () => {
    const balance = await get_balance('B508ED0D54597D516A680E7951F18CAD24C7EC9F');
    element = document.getElementById("get_balance");

    echo();
    echo("balance:");
    echo(balance);
}

document.getElementById("send_coin_btn").onclick = async () => {
    await send_coin('B508ED0D54597D516A680E7951F18CAD24C7EC9F', '0D9FE6A785C830D2BE66FE40E0E7FE3D9838456C', '88', 'wmyZZoMedWlsPUDVCOy+TiVcrIBPcn3WJN8k5cPQgIvC8cbcR10FtdAdzIlqXQJL9hBw1i0RsVjF6Oep/06Ezg==', 'wvHG3EddBbXQHcyJal0CS/YQcNYtEbFYxejnqf9OhM4=');

    const balance = await get_balance('B508ED0D54597D516A680E7951F18CAD24C7EC9F');
    element = document.getElementById("send_coin");

    echo();
    echo("new balance:");
    echo(balance);

}