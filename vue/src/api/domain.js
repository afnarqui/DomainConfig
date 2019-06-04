export default function buscardomain (name) {
  return fetch(name)
     .then(response => response.json())
     .then(json => {
       return json
      }
       )
}
