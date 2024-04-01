<script>
import InvoiceForm  from './InvoiceForm.vue'

export default {
  components:{
    InvoiceForm
  },
  props: [
    "title",
    "dataSetupForm"
  ],
  emits:[
    "my_submit1"
  ]
}

</script>

<template>
  <h1>{{ title }}</h1>

  <p>{{ dataSetupForm }}</p>

  <InvoiceForm @my_submit1="(form) => {
    console.log('By now, the form should\'ve been set.');
    console.log('Form issuer: '+form.issuer);
  }" />

</template>