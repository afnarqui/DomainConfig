<template>
<div id="id">
  <form id="main" v-cloak>

    <div class="bar">
         <br/>
        <input type="text" v-model="busc" placeholder="Enter your Search" />
        <br/>
        <br/>
        <input @click="buscar"  type="button" value="Search" class="btn btn-success">
        <br/>
        <br/>
    </div>

 <b-card-group deck v-if="posts && postsnew.length">
  <b-card header="Data">
    <b-list-group v-for="post of postsnew" v-bind:key="post.id">
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
    <b-card header="Endpoints">
    <ul >
        <li v-for="po of posts.endpoints" v-bind:key="po.id">
           <h4>Server Name: {{po.serverName}}</h4>
           <p>Status Message: {{po.statusMessage}}</p>
           <p>Ip Address: {{po.ipAddress}}</p>
          <p><strong>Address: {{po.address}}</strong></p>
          <p>Grade: {{po.grade}}</p>
          <p><strong>Status Message: {{po.statusMessage}}</strong></p>
          <p><strong>Grade Trust Ignored: {{po.gradeTrustIgnored}}</strong></p>
          <p><strong>Has Warnings: {{po.hasWarnings}}</strong></p>
          <p><strong>Is Exceptional: {{po.isExceptional}}</strong></p>
          <p><strong>Progress: {{po.progress}}</strong></p>
          <p><strong>Duration: {{po.duration}}</strong></p>
          <p><strong>Delegation: {{po.delegation}}</strong></p>
        </li>
    </ul>
  </b-card>

</b-card-group>

</form>

</div>


</template>
<script>
import getDominios from '../api/info'

export default {
  name: 'Domain',
  data () {
    return {
      posts: [],
      postsnew: [],
    }
  },
  mounted() {
  },
  methods: {
    buscar: function () {
        const self = this
        const url = `http://localhost:8001/public?name=${self.busc}`
        getDominios(url)
        .then( function(dominios) {
          let data = JSON.stringify(dominios)
          let datanew = `[${data}]`
          self.postsnew = JSON.parse(datanew)
          self.posts = JSON.parse(data)
        })
    .catch((e) => {
      console.error(e)
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
