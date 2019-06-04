<template>
  <b-navbar
    toggleable="md"
    type="dark"
    variant="info">
    <b-navbar-brand :to="{ name: 'Domainall' }">Domain All</b-navbar-brand>
    <b-navbar-nav>
      <b-navbar-brand :to="{ name: 'Domain' }">Domain</b-navbar-brand>
      <b-nav-item :to="{ name: 'Domaincompare' }">Domain compare</b-nav-item>
    </b-navbar-nav>
  </b-navbar>
</template>
<script>
export default {
  name: 'Navbar'
};
</script>
