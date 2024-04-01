<script>
export default {
    data(){
        return {
            form: {
                issuer: "",
                counterpart: "",
                shutup: "",
            }
        }
    },
    emits:[
        'my_submit1'
    ],
    methods: {
        submit(){
            console.log("form-submit called!") //emission not called, pick up from here
            console.log([
                this.form.issuer,
                this.form.counterpart,
                this.form.shutup
            ])
            this.$emit('my_submit1', this.form)
        }
    }
}
</script>

<template>
    <form>
        <label for="issuer">Issuer: </label>
        <input name="issuer" v-model="form.issuer" /><br>

        <label for="counterpart">Counterpart: </label>
        <input name="counterpart" v-model="form.counterpart" /><br>

        <label for="shutup">Shut up: </label>
        <input name="shutup" v-model="form.shutup" /><br>

    </form>
    
    <button @click="submit">Submit</button>

    <p>{{ form.issuer }}</p><br>
    <p>{{ form.counterpart }}</p><br>
    <p>{{ form.shutup }}</p><br>
</template>