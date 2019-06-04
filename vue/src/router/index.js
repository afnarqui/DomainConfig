import Vue from 'vue';
import Router from 'vue-router';
import Domain from '@/components/Domain';

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/Domain',
      name: 'Domain',
      component: Domain
    }
  ]
});
