<script setup>
import { Fancybox } from '@fancyapps/ui/dist/index.esm.js';

const props = defineProps({
  options: Object,
});
const container = ref(null);

onMounted(() => {
  Array.from(container.value.children).map((el) => {
    el.setAttribute('data-fancybox', 'gallery');
  });
  Fancybox.bind(container.value, '[data-fancybox]', {
    ...(props.options || {}),
  });
});

nextTick(() => {
  Fancybox.unbind(container.value);
  Fancybox.close();

  Fancybox.bind(container.value, '[data-fancybox]', {
    ...(props.options || {}),
  });
});

onUnmounted(() => {
  Fancybox.destroy();
});

</script>

<template>
  <div ref="container">
    <slot></slot>
  </div>
</template>

<style></style>
