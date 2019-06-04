
export default function getDominios (name) {
  return fetch(name)
     .then(response => response.json())
     .then(json => {
       return json
      }
       )
}
