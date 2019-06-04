<template>
<div id="id">
 <div class="bar">
         <br/>

         <b-card-group deck v-if="postdomain.length">
        <b-card header="Domain all">
          <b-list-group v-for="post of postdomain" v-bind:key="post.id">
            <b-list-group-item href="#"> Host: {{post.host}} Port: {{post.port}}
            </b-list-group-item>
          </b-list-group>

      </b-card-group>
    </div>
</div>
</template>
<script>
import buscardomain from '../api/domain'
import path from '../api/config'

export default {
  name: 'Domainall',
  data () {
    return {
      postdomain: [],
    }
  },
  mounted() {
     this.buscarDomains()
  },
  methods: {
    buscarDomains: function () {
        const self = this
        let url = "http://localhost:8001/searchdomain"
    buscardomain(url)
    .then( function(buscar){
          self.postdomain = buscar
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
