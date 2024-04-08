<script setup>
import { Fancybox } from '@fancyapps/ui';
import '@fancyapps/ui/dist/fancybox/fancybox.css';

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

onUpdated(() => {
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
