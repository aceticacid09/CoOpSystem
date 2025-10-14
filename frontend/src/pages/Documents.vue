<template>
  <div>
    <Header />
    <div class="container">
      <Breadcrumb :items="[{ label: 'เอกสาร', link: '/documents' }]" />
      <DocumentsSection />
    </div>
    <BackToTop />
    <Footer />
  </div>
</template>

<script setup>
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import BackToTop from "../components/BackToTop.vue";
import Breadcrumb from "../components/Breadcrumb.vue";
import DocumentsSection from "../components/DocumentsSections.vue";
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;700&display=swap");

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
