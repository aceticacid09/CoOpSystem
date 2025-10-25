<!-- frontend/src/components/DateRangeDropdown.vue -->
<template>
  <div class="filter-group">
    <label class="filter-label">{{ label }}</label>
    <div class="date-filter-dropdown" ref="dateFilterRef">
      <button class="date-dropdown-button" @click.stop="toggleDropdown">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
          stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
          <line x1="16" y1="2" x2="16" y2="6" />
          <line x1="8" y1="2" x2="8" y2="6" />
          <line x1="3" y1="10" x2="21" y2="10" />
        </svg>
        <span>{{ dateRangeDisplay }}</span>
        <svg class="dropdown-arrow" :class="{ open: isOpen }" xmlns="http://www.w3.org/2000/svg" width="16" height="16"
          viewBox="0 0 24 24" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
          stroke-linejoin="round">
          <polyline points="6 9 12 15 18 9"></polyline>
        </svg>
      </button>
      <transition name="dropdown-fade">
        <div v-if="isOpen" class="date-dropdown-menu">
          <div class="date-inputs">
            <div class="input-group">
              <label>จาก</label>
              <input type="date" v-model="localDateFrom" class="date-input" />
            </div>
            <div class="input-group">
              <label>ถึง</label>
              <input type="date" v-model="localDateTo" class="date-input" />
            </div>
          </div>
          <div class="date-actions">
            <button class="btn-clear-date" @click="handleClear">ล้าง</button>
            <button class="btn-apply-date" @click="handleApply">ใช้งาน</button>
          </div>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue';

const props = defineProps({
  label: {
    type: String,
    default: 'วันที่'
  },
  dateFrom: {
    type: String,
    default: ''
  },
  dateTo: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: 'เลือกช่วงวันที่'
  }
});

const emit = defineEmits(['update:dateFrom', 'update:dateTo', 'apply', 'clear']);

const isOpen = ref(false);
const dateFilterRef = ref(null);
const localDateFrom = ref(props.dateFrom);
const localDateTo = ref(props.dateTo);

const dateRangeDisplay = computed(() => {
  if (localDateFrom.value && localDateTo.value) {
    return `${formatDateShort(localDateFrom.value)} - ${formatDateShort(localDateTo.value)}`;
  }
  return props.placeholder;
});

const formatDateShort = (dateStr) => {
  if (!dateStr) return '';
  const d = new Date(dateStr);
  if (isNaN(d)) return '';
  const day = d.getDate();
  const month = d.getMonth() + 1;
  const year = d.getFullYear() + 543;
  return `${day}/${month}/${year}`;
};

const toggleDropdown = () => {
  isOpen.value = !isOpen.value;
};

const handleApply = () => {
  emit('update:dateFrom', localDateFrom.value);
  emit('update:dateTo', localDateTo.value);
  emit('apply', {
    dateFrom: localDateFrom.value,
    dateTo: localDateTo.value
  });
  isOpen.value = false;
};

const handleClear = () => {
  localDateFrom.value = '';
  localDateTo.value = '';
  emit('update:dateFrom', '');
  emit('update:dateTo', '');
  emit('clear');
};

const handleClickOutside = (event) => {
  if (isOpen.value && dateFilterRef.value && !dateFilterRef.value.contains(event.target)) {
    isOpen.value = false;
  }
};

watch(() => props.dateFrom, (newVal) => {
  localDateFrom.value = newVal;
});

watch(() => props.dateTo, (newVal) => {
  localDateTo.value = newVal;
});

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
});

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside);
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

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

.date-filter-dropdown {
  position: relative;
  width: 400px;
  margin-bottom: 10px;
}

.date-dropdown-button {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  width: 100%;
  padding: 9px 14px;
  background: #fff;
  border: 1px solid #e5e7e6;
  border-radius: 10px;
  font-size: 14px;
  color: #333;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  box-shadow: 0 1px 3px rgba(10, 10, 10, 0.02);
}

.date-dropdown-button:hover {
  border-color: #037266;
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.08);
}

.date-dropdown-button svg:first-child {
  flex-shrink: 0;
  color: #666;
}

.date-dropdown-button span {
  flex: 1;
  text-align: left;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.dropdown-arrow {
  transition: transform 0.25s ease;
  flex-shrink: 0;
  margin-left: 8px;
}

.dropdown-arrow.open {
  transform: rotate(180deg);
}

.date-dropdown-menu {
  position: absolute;
  top: calc(100% + 6px);
  left: 0;
  width: 100%;
  background: #fff;
  border: 1px solid #e5e7e6;
  border-radius: 10px;
  box-shadow: 0 8px 24px rgba(6, 20, 18, 0.12);
  z-index: 100;
  padding: 12px;
}

.date-inputs {
  display: flex;
  gap: 12px;
  margin-bottom: 12px;
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
  flex: 1;
}

.input-group label {
  font-size: 13px;
  color: #666;
  font-weight: 600;
}

.date-input {
  padding: 10px 12px;
  border: 2px solid #e5e5e5;
  border-radius: 8px;
  font-size: 14px;
  font-family: "Kanit", sans-serif;
  transition: all 0.2s;
}

.date-input:focus {
  outline: none;
  border-color: #037266;
}

.date-actions {
  display: flex;
  gap: 8px;
  padding-top: 12px;
  border-top: 1px solid #f0f0f0;
}

.btn-clear-date,
.btn-apply-date {
  flex: 1;
  padding: 8px 16px;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-clear-date {
  background: #f0f0f0;
  color: #666;
}

.btn-clear-date:hover {
  background: #e0e0e0;
}

.btn-apply-date {
  background: #037266;
  color: #fff;
}

.btn-apply-date:hover {
  background: #026b5b;
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

@media (max-width: 768px) {
  .date-filter-dropdown {
    width: 100%;
  }
  
  .date-inputs {
    flex-direction: column;
  }
}
</style>