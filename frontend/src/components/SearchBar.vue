<!-- frontend/src/components/SearchBar.vue -->
<template>
  <div class="search-sort">
    <div class="search-box">
      <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="11" cy="11" r="8"></circle>
        <path d="m21 21-4.35-4.35"></path>
      </svg>
      <input type="text" :value="modelValue" @input="$emit('update:modelValue', $event.target.value)"
        :placeholder="placeholder" />
    </div>
    <button class="btn-search" @click="$emit('search')">ค้นหา</button>
    <button @click="$emit('toggle-sort')" class="btn-sort" :title="isAscending ? 'เรียงจากเก่าสุด' : 'เรียงจากใหม่สุด'">
      <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" fill="none" stroke="#9CA3AF" stroke-width="2"
        stroke-linecap="round" stroke-linejoin="round">
        <path d="M3 6h18M3 12h12M3 18h6" />
      </svg>
    </button>
  </div>
</template>

<script setup>
defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: 'ค้นหา...'
  },
  isAscending: {
    type: Boolean,
    default: true
  }
});

defineEmits(['update:modelValue', 'search', 'toggle-sort']);
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;700&display=swap");

* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.search-sort {
  display: flex;
  align-items: center;
  margin-top: 10px;
  margin-bottom: 10px;
  flex-wrap: wrap;
  gap: 12px;
}

.search-box {
  flex: 1;
  min-width: 250px;
  position: relative;
  display: flex;
  align-items: center;
}

.search-box svg {
  position: absolute;
  left: 14px;
  color: #999;
}

.search-box input {
  width: 100%;
  padding: 11px 14px 11px 42px;
  border: 2px solid #e5e5e5;
  border-radius: 10px;
  font-size: 14px;
  transition: all 0.2s;
  font-family: "Kanit", sans-serif;
}

.search-box input:focus {
  outline: none;
  border-color: #037266;
}

.btn-search {
  background: #037266;
  color: #fff;
  border: none;
  padding: 9px 18px;
  border-radius: 10px;
  cursor: pointer;
  font-weight: 600;
  box-shadow: 0 6px 18px rgba(3, 114, 102, 0.08);
  font-size: 14px;
  transition: all 0.2s;
  height: 45px;
  width: 90px;
}

.btn-search:hover {
  transform: translateY(-2px);
}

.btn-sort {
  background: #f1f1f1;
  border: none;
  padding: 8px 10px;
  border-radius: 10px;
  cursor: pointer;
  height: 45px;
  width: 70px;
}

.btn-sort:hover {
  background: #e5e5e5;
}

/* Responsive */
@media (max-width: 650px) {
  .search-sort {
    margin: 20px 0;
  }

  .search-box {
    min-width: 200px;
  }
}
</style>