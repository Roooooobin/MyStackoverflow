export function getUserData(uid){
    return fetch(`http://0.0.0.0:8080/user/get?uid=${uid}`)
    .then(data=>data.json())
    .then(data=>data)
}