<!-- frontend/src/pages/NewsEvents.vue -->
<template>
  <div>
    <Header />
    <div class="container">
      <Breadcrumb :items="[{ label: 'ข่าวสารและกิจกรรม', link: '/news' }]" />
      <NewsSection />
    </div>
    <BackToTop />
    <Footer />
  </div>
</template>

<script setup>
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import NewsSection from "../components/NewsSections.vue";
import BackToTop from "../components/BackToTop.vue";
import Breadcrumb from "../components/Breadcrumb.vue";
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;700&display=swap");

/* ========================================
   GLOBAL STYLES
======================================== */
* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.container {
  max-width: 1200px;
  margin: auto;
  padding: 0 20px;
}
</style>