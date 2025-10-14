<template>
  <button
    v-show="showButton"
    class="back-to-top"
    @click="scrollToTop"
    aria-label="Scroll to top"
  >
    <!-- SVG ลูกศรขึ้น -->
    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <polyline points="18 15 12 9 6 15"></polyline>
    </svg>
  </button>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from "vue";

const showButton = ref(false);

const handleScroll = () => {
  showButton.value = window.scrollY > 0;
};

const scrollToTop = () => {
  window.scrollTo({ top: 0, behavior: "smooth" });
};

onMounted(() => window.addEventListener("scroll", handleScroll));
onUnmounted(() => window.removeEventListener("scroll", handleScroll));
</script>

<style scoped>
.back-to-top {
  position: fixed;
  bottom: 30px;
  right: 30px;
  width: 48px;
  height: 48px;
  background-color: #037266; /* สีหลัก */
  color: #ffffff;
  border: none;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.15);
  z-index: 9999;
  opacity: 0.9;
  transition: opacity 0.3s, transform 0.3s, background-color 0.3s;
}

.back-to-top:hover {
  background-color: #025a50;
  opacity: 1;
  transform: translateY(-2px) scale(1.05);
}

.back-to-top svg {
  stroke: #ffffff;
}
</style>
