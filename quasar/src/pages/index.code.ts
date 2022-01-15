import { ref, onMounted } from 'vue';
import { Loading, QSpinnerGears } from 'quasar';

const useHellWorld = () => {

    const data = ref<Array<string>>([
        'hell',
        'world'
    ]);
    // eslint-disable-next-line prefer-const
    let pageInit = ref<boolean>(false);

    const deleteBtn = (idx: number) => {
        console.log('idx', idx);
        data.value.splice(idx, 1);
    }

    onMounted(() => {
        Loading.show({
        spinner: QSpinnerGears
        })

        setTimeout(() => {
        data.value.push('asdfasdfa')
        data.value.push('asdfasdfa')
        data.value.push('asdfasdfa')

        Loading.hide()

        pageInit.value = true;
        }, 2000)
    })

    return {
        deleteBtn,
        data,
        pageInit
    }
}

export {
    useHellWorld
}