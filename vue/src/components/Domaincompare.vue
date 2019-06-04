<template>
<div id="id">
   <div class="bar">

        <br/><br/>
 <b-card-group deck v-if="posts && postdomaincomparar.length">
  <b-card header="Domain">
    <b-list-group v-for="post of postdomaincomparar[0]" v-bind:key="post.id">
      <b-list-group-item href="#"> Host: {{post.host}}</b-list-group-item>
      <b-list-group-item href="#">Port: {{post.port}}</b-list-group-item>
      <b-list-group-item href="#">Is Public: {{post.isPublic}}</b-list-group-item>
      <b-list-group-item href="#">Status: {{post.status}}</b-list-group-item>
      <b-list-group-item href="#">Start Time: {{post.startTime}}</b-list-group-item>
      <b-list-group-item href="#">Test Time: {{post.testTime}}</b-list-group-item>
      <b-list-group-item href="#">Engine Version: {{post.engineVersion}}</b-list-group-item>
      <b-list-group-item href="#">Criteria Version: {{post.criteriaVersion}}</b-list-group-item>
       <p class="card-text mt-2">
        Protocol:  {{post.protocol}}
      </p>
    </b-list-group>
  </b-card>
    <b-card header="Domain compare">
        <b-list-group v-for="post of postdomaincomparar[1]" v-bind:key="post.id">
           <div v-if="post.host.includes('Valor Editado:') === true">
              <b-list-group-item style="background-color:blue;" href="#"> Host: {{post.host}}</b-list-group-item>
          </div>
          <div v-else-if="post.host.includes('Valor Agregado:') === true">
             <b-list-group-item style="background-color:#405d27;" href="#"> Host: {{post.host}}</b-list-group-item>
          </div>
          <div v-else-if="post.host.includes('Valor Eliminado:') === true">
             <b-list-group-item style="background-color:#034f84;" href="#"> Host: {{post.host}}</b-list-group-item>
          </div>
          <div v-else>
            <b-list-group-item href="#"> Host: {{post.host}}</b-list-group-item>
          </div>

           <div v-if="post.port.includes('Valor Editado:') === true">
              <b-list-group-item style="background-color:blue;" href="#"> port: {{post.port}}</b-list-group-item>
          </div>
          <div v-else-if="post.port.includes('Valor Agregado:') === true">
             <b-list-group-item style="background-color:#405d27;" href="#"> port: {{post.port}}</b-list-group-item>
          </div>
          <div v-else-if="post.port.includes('Valor Eliminado:') === true">
             <b-list-group-item style="background-color:#034f84;" href="#"> port: {{post.port}}</b-list-group-item>
          </div>
          <div v-else>
            <b-list-group-item href="#"> port: {{post.port}}</b-list-group-item>
          </div>

           <div v-if="post.status.includes('Valor Editado:') === true">
              <b-list-group-item style="background-color:blue;" href="#"> status: {{post.status}}</b-list-group-item>
          </div>
          <div v-else-if="post.status.includes('Valor Agregado:') === true">
             <b-list-group-item style="background-color:#405d27;" href="#"> status: {{post.status}}</b-list-group-item>
          </div>
          <div v-else-if="post.status.includes('Valor Eliminado:') === true">
             <b-list-group-item style="background-color:#034f84;" href="#"> status: {{post.status}}</b-list-group-item>
          </div>
          <div v-else>
            <b-list-group-item href="#"> status: {{post.status}}</b-list-group-item>
          </div>

           <div v-if="post.ispublic.includes('Valor Editado:') === true">
              <b-list-group-item style="background-color:blue;" href="#"> ispublic: {{post.ispublic}}</b-list-group-item>
          </div>
          <div v-else-if="post.ispublic.includes('Valor Agregado:') === true">
             <b-list-group-item style="background-color:#405d27;" href="#"> ispublic: {{post.ispublic}}</b-list-group-item>
          </div>
          <div v-else-if="post.ispublic.includes('Valor Eliminado:') === true">
             <b-list-group-item style="background-color:#034f84;" href="#"> ispublic: {{post.ispublic}}</b-list-group-item>
          </div>
          <div v-else>
            <b-list-group-item href="#"> ispublic: {{post.ispublic}}</b-list-group-item>
          </div>

           <div v-if="post.protocol.includes('Valor Editado:') === true">
              <b-list-group-item style="background-color:blue;" href="#"> protocol: {{post.protocol}}</b-list-group-item>
          </div>
          <div v-else-if="post.protocol.includes('Valor Agregado:') === true">
             <b-list-group-item style="background-color:#405d27;" href="#"> protocol: {{post.protocol}}</b-list-group-item>
          </div>
          <div v-else-if="post.protocol.includes('Valor Eliminado:') === true">
             <b-list-group-item style="background-color:#034f84;" href="#"> protocol: {{post.protocol}}</b-list-group-item>
          </div>
          <div v-else>
            <b-list-group-item href="#"> protocol: {{post.protocol}}</b-list-group-item>
          </div>

       <p class="card-text mt-2">
        Protocol:  {{post.protocol}}
      </p>
    </b-list-group>
  </b-card>

</b-card-group>
    </div>

</div>


</template>
<script>
import getDominios from '../api/info'
import buscardomain from '../api/domain'
import buscardomaincomparar from '../api/domaincomparar'

export default {
  name: 'Domaincompare',
  data () {
    return {
      posts: [],
      postsnew: [],
      postdomain: [],
      postdomaincomparar: [],
      color: 'blue',
    }
  },
  mounted() {
    this.buscarDomainComparars()
  },
  methods: {
    buscar: function () {
        const self = this
        const url = `http://localhost:8001/public?nombre=${self.busc}`
        getDominios(url)
        .then( function(dominios) {
          let data = JSON.stringify(dominios)
          let datanew = `[${data}]`
          self.postsnew = JSON.parse(datanew)
          console.log(self.postsnew)
          console.log(JSON.parse(data))
          self.posts = JSON.parse(data)
        })
    .catch((e) => {
      console.error(e)
    })
    },
    buscarDomains: function () {
        const self = this
        const url = `http://localhost:8001/buscardomain`
    buscardomain(url)
    .then( function(buscar){
          self.postdomain = buscar

    }).catch((e)=> {
      console.log(e)
    })
    },
    buscarDomainComparars: function () {
      const self = this
      const url = `http://localhost:8001/searchdomaincomparar`
    buscardomaincomparar(url)
    .then( function(buscar){
          self.postdomaincomparar = buscar
    }).catch((e)=> {
      console.log(e)
    })

    }
  }
}
</script>
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
