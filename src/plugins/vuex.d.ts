import { ComponentCustomProperties } from 'vue'
import { Store } from 'vuex'

declare module '@vue/runtime-core' {
    interface State {
        count: number
    }

    // 为 `this.$store` 提供类型声明
    interface ComponentCustomProperties {
        $store: Store<State>
    }
}