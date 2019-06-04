import Vue from 'vue';
import Router from 'vue-router';
import Domain from '@/components/Domain';
import Domainall from '@/components/Domainall';

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/Domain',
      name: 'Domain',
      component: Domain
    },
    {
      path: '/',
      name: 'Domainall',
      component: Domainall
    }
  ]
});
