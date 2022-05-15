export function getTopic() {
    return fetch(`http://0.0.0.0:8080/topic/list`)
      .then(data => data.json())
      .then(data => data.data)
  }