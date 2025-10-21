<!-- frontend/src/components/CustomDropdown.vue -->
<template>
  <div class="filter-group">
    <label v-if="label" class="filter-label">{{ label }}</label>
    <div class="custom-dropdown" ref="dropdownRef" @click="toggleDropdown">
      <div class="dropdown-selected">
        <span>{{ modelValue }}</span>
        <svg class="dropdown-arrow" :class="{ open: isOpen }" xmlns="http://www.w3.org/2000/svg" width="16" height="16"
          viewBox="0 0 24 24" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
          stroke-linejoin="round">
          <polyline points="6 9 12 15 18 9"></polyline>
        </svg>
      </div>
      <transition name="dropdown-fade">
        <ul v-if="isOpen" class="dropdown-menu">
          <li v-for="option in options" :key="option" @click.stop="selectOption(option)"
            :class="{ selected: modelValue === option }">
            {{ option }}
          </li>
        </ul>
      </transition>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';

const props = defineProps({
  label: {
    type: String,
    default: ''
  },
  modelValue: {
    type: [String, Number],
    required: true
  },
  options: {
    type: Array,
    required: true
  },
  width: {
    type: String,
    default: '300px'
  }
});

const emit = defineEmits(['update:modelValue']);

const isOpen = ref(false);
const dropdownRef = ref(null);

const toggleDropdown = () => {
  isOpen.value = !isOpen.value;
};

const selectOption = (option) => {
  emit('update:modelValue', option);
  isOpen.value = false;
};

const handleClickOutside = (event) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target)) {
    isOpen.value = false;
  }
};

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
});

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside);
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;600;700&display=swap");

* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.filter-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.filter-label {
  font-size: 13px;
  color: #6b6b6b;
  font-weight: 500;
}

.custom-dropdown {
  position: relative;
  width: v-bind(width);
  cursor: pointer;
  user-select: none;
  margin-bottom: 10px;
}

.dropdown-selected {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 9px 14px;
  background: #fff;
  border: 1px solid #e5e7e6;
  border-radius: 10px;
  font-size: 14px;
  color: #333;
  font-weight: 500;
  transition: all 0.2s ease;
  box-shadow: 0 1px 3px rgba(10, 10, 10, 0.02);
}

.dropdown-selected:hover {
  border-color: #037266;
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.08);
}

.dropdown-arrow {
  transition: transform 0.25s ease;
  flex-shrink: 0;
  margin-left: 8px;
}

.dropdown-arrow.open {
  transform: rotate(180deg);
}

.dropdown-menu {
  position: absolute;
  top: calc(100% + 6px);
  left: 0;
  width: 100%;
  background: #fff;
  border: 1px solid #e5e7e6;
  border-radius: 10px;
  box-shadow: 0 8px 24px rgba(6, 20, 18, 0.12);
  max-height: 240px;
  overflow-y: auto;
  z-index: 100;
  list-style: none;
  margin: 0;
  padding: 6px;
}

.dropdown-menu li {
  padding: 9px 12px;
  font-size: 14px;
  color: #333;
  cursor: pointer;
  border-radius: 6px;
  transition: all 0.15s ease;
  font-weight: 500;
  white-space: nowrap;
}

.dropdown-menu li:hover {
  background: #f0f8f7;
  color: #037266;
}

.dropdown-menu li.selected {
  background: #e6f4f2;
  color: #037266;
  font-weight: 600;
}

.dropdown-fade-enter-active,
.dropdown-fade-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
  transform-origin: top;
}

.dropdown-fade-enter-from {
  opacity: 0;
  transform: translateY(-8px) scale(0.95);
}

.dropdown-fade-leave-to {
  opacity: 0;
  transform: translateY(-4px) scale(0.98);
}

/* Responsive */
@media (max-width: 900px) {
  .custom-dropdown {
    width: 200px;
  }
}

@media (max-width: 650px) {
  .custom-dropdown {
    width: 100%;
  }
}
</style>