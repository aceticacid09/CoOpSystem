<!-- frontend/src/components/SegmentedControl.vue -->
<template>
    <div class="filter-group">
        <label v-if="label" class="filter-label">{{ label }}</label>
        <div class="segmented-control">
            <button v-for="option in options" :key="option" :class="{ active: modelValue === option }"
                @click="selectOption(option)">
                {{ option }}
            </button>
        </div>
    </div>
</template>

<script setup>
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
    }
});

const emit = defineEmits(['update:modelValue']);

const selectOption = (option) => {
    emit('update:modelValue', option);
};
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

.segmented-control {
    display: inline-flex;
    border-radius: 12px;
    background: #f7f9f8;
    padding: 4px;
    gap: 4px;
    border: 1px solid #eef4f2;
    box-shadow: 0 1px 3px rgba(10, 10, 10, 0.02);
    margin-bottom: 7px;
}

.segmented-control button {
    min-width: 44px;
    padding: 8px 12px;
    border-radius: 8px;
    border: none;
    background: transparent;
    cursor: pointer;
    font-weight: 600;
    color: #556;
    transition: background .18s, color .18s, transform .08s;
}

.segmented-control button.active {
    background: linear-gradient(180deg, #037266, #026b5b);
    color: #fff;
    box-shadow: 0 6px 16px rgba(3, 114, 102, 0.12);
    transform: translateY(-1px);
}

/* Responsive */
@media (max-width: 650px) {
    .segmented-control {
        width: 100%;
        justify-content: space-between;
    }
}
</style>