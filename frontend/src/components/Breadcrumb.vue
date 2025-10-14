<template>
  <nav class="breadcrumb">
    <ul>
      <li>
        <a href="/" class="home-link">
          <!-- บ้าน outline แบบ minimal -->
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="icon-home">
            <path d="M3 12l9-9 9 9M4 10v10h6v-6h4v6h6V10" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" fill="none"/>
          </svg>
          หน้าแรก
        </a>
      </li>
      <li v-for="(item, index) in items" :key="index">
        <span class="separator">›</span>
        <template v-if="index !== items.length - 1">
          <a :href="item.link">{{ item.label }}</a>
        </template>
        <template v-else>
          <span class="current">{{ item.label }}</span>
        </template>
      </li>
    </ul>
  </nav>
</template>

<script setup>
defineProps({
  items: {
    type: Array,
    required: true
  }
});
</script>

<style scoped>

@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;700&display=swap");

* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.breadcrumb {
  font-size: 0.95rem;
  margin: 50px 0 50px 20px; /*margin: [บน] [ขวา] [ล่าง] [ซ้าย]; */
  font-weight: 500;
  color: #555;
}

.breadcrumb ul {
  list-style: none;
  display: flex;
  align-items: center;
  padding: 0;
  margin: 0;
}

.breadcrumb li {
  display: flex;
  align-items: center;
}

.home-link {
  display: flex;
  align-items: center;
  color: #444;
  font-weight: 600;
  text-decoration: none;
  transition: color 0.2s;
}

.home-link:hover {
  color: #45805f; /* สีทองอ่อน หรูหรา */
}

.icon-home {
  width: 18px;
  height: 18px;
  margin-right: 6px;
  stroke: #444;
  transition: stroke 0.2s;
}

.home-link:hover .icon-home {
  stroke: #45805f;
}

.separator {
  margin: 0 8px;
  color: #ccc;
}

.breadcrumb a {
  text-decoration: none;
  color: #555;
  transition: color 0.2s;
}

.breadcrumb a:hover {
  color: #45805f;
}

.current {
  color: #333;
  font-weight: 600;
}
</style>
