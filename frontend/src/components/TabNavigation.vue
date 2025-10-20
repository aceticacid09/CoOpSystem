<!-- frontend/src/components/TabNavigation.vue -->
<template>
  <div class="tabs">
    <button v-for="tab in tabs" :key="tab" class="tab-btn" :class="{ active: modelValue === tab }" @click="selectTab(tab)">
      {{ tab }}
    </button>
  </div>
</template>

<script setup>
const props = defineProps({
  tabs: {
    type: Array,
    required: true
  },
  modelValue: {
    type: [String, Number],
    required: true
  }
});

const emit = defineEmits(['update:modelValue']);

const selectTab = (tab) => {
  emit('update:modelValue', tab);
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;700&display=swap");

* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.tabs {
  display: flex;
  gap: 25px;
  border-bottom: 1px solid #eee;
  margin-bottom: 20px;
  flex-wrap: wrap;
}

.tab-btn {
  background: none;
  border: none;
  font-size: 16px;
  color: #555;
  padding-bottom: 8px;
  cursor: pointer;
  position: relative;
  transition: 0.2s;
}

.tab-btn.active {
  color: #037266;
  font-weight: 600;
}

.tab-btn.active::after {
  content: "";
  position: absolute;
  left: 0;
  bottom: -1px;
  width: 100%;
  height: 2px;
  background: #037266;
}
</style>