<template>
  <TableContainer :length="sortedEnrichedValidators.length" :columns="properties" :sort="sort" :loaded="loaded">
    <ValidatorRow
      v-for="(validator, index) in sortedEnrichedValidators"
      :key="validator.operatorAddress"
      :index="index"
      :validator="validator"
      :delegation="getDelegation(validator)"
      :rewards="getRewards(validator)"
      :staking-denom="stakingDenom"
    />
    <template v-slot:empty>
      <slot name="empty"></slot>
    </template>
  </TableContainer>
</template>

<script>
import { orderBy } from 'lodash'
import TableContainer from '@/components/TableContainer'
import ValidatorRow from '@/components/ValidatorRow'

export default {
  name: `TableValidators`,
  components: { ValidatorRow, TableContainer },
  props: {
    validators: {
      type: Array,
      required: true,
    },
    delegations: {
      type: Array,
      default: () => [],
    },
    rewards: {
      type: Array,
      default: () => [],
    },
    searchTerm: {
      type: Boolean,
      default: false,
    },
    loaded: {
      type: Boolean,
      default: true,
    },
  },
  data: () => ({
    sort: {
      property: `votingPower`,
      order: `desc`,
    },
    stakingDenom: 'cgas',
  }),
  computed: {
    sortedEnrichedValidators() {
      const orderedValidators = orderBy(
        this.validators.map((validator) => ({
          ...validator,
          smallName: validator.name ? validator.name.toLowerCase() : '',
        })),
        [this.sort.property],
        [this.sort.order],
      )
      return orderedValidators
    },
    properties() {
      return [
        {
          title: `Status`,
          value: `status`,
        },
        {
          title: `Name`,
          value: `smallName`,
        },
        {
          title: `Rewards`,
          value: `expectedReturns`,
        },
        {
          title: `Voting Power`,
          value: `votingPower`,
        },
      ]
    },
  },
  methods: {
    getDelegation({ operatorAddress }) {
      return this.delegations.find(({ validator }) => validator.operatorAddress === operatorAddress)
    },
    getRewards({ operatorAddress }) {
      if (this.rewards) {
        return (
          this.rewards
            /* istanbul ignore next */
            .filter(({ validator }) => validator.operatorAddress === operatorAddress)
        )
      }
    },
  },
}
</script>
<style scoped>
.no-results {
  text-align: center;
  margin: 3rem;
  color: #4a5568;
  font-size: 12px;
}

.sortingOptions {
  margin: 0.5rem 1rem;
}

.sortingOptions li.active {
  color: hsl(7, 88%, 60%);
}

.sortingOptions li {
  padding: 1rem 0.5rem;
  cursor: pointer;
  display: flex;
  justify-content: space-between;
}

.sorting-check {
  justify-content: space-between;
}

.sorting-check.inactive {
  /*color: var(--app-bg);*/
}

.sortingOptions .material-icons {
  font-size: 22px;
  width: 2rem;
  vertical-align: text-bottom;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}

.fade-enter,
.fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}

@media screen and (min-width: 768px) {
  .sortingOptions {
    display: none;
  }
}
</style>
