async function de(txt) {
    const passphrase = document.getElementById("password").value
    const raw = Uint8Array.from(atob(txt), (m) => m.codePointAt(0));
    return new TextDecoder().decode(await window.crypto.subtle.decrypt({name: 'AES-GCM', iv: raw.slice(0, 12),}, await window.crypto.subtle.importKey("raw", new TextEncoder().encode(passphrase), {"name": "AES-GCM"}, false, ["encrypt", "decrypt"]), raw.slice(12)))
}

async function load() {
    let promises = []
    data.a.forEach((v) => {
        promises.push(
            de(v).then((v) => {
                document.body.insertAdjacentHTML('beforeend', v);
            }).catch(() => {
                return Promise.reject("s")
            })
        )
    })
    data.b.forEach((v) => {
        promises.push(
            de(v).then((v) => {
                document.head.insertAdjacentHTML("beforeend", '<style>'+v+'</style>')
            })
        )
    })
    data.c.forEach((v) => {
        promises.push(
            de(v).then((v) => {
                const el = document.createElement('script')
                el.text = v
                document.body.appendChild(el);
            })
        )
    })
    Promise.all(promises)
        .then(() => {
            document.getElementById("aewef").remove()
            document.getElementById("asdsd").remove()
            document.getElementById("jdjdwjsh").remove()
        }).catch(() => {
    })
}

document.getElementById('form').addEventListener('submit', async (event) => {
    event.preventDefault()
    await load()
})