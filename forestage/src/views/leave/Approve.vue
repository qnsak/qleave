<script setup>
  import { ref, reactive, onMounted } from 'vue'
  import api from "@/api"

  onMounted(() => {
    getLeaveList()
  })

  const pageNow = ref(0)
  const leaveList = reactive([])
  const pagingInfo = reactive({
    from: 0,
    to: 0,
    total: 0,
    prev: 0,
    next: 0,
  })

  /**
   * 取得請假申請
   */
  const getLeaveList = async (page = 0) => {
      pageNow.value = page;
      let getLeaveListResult = await api.leave.getApproveLeaveList(page)
      leaveList.length = 0;
      Object.assign(leaveList, getLeaveListResult.data);
      let prevUrl = 0;
      let nextUrl = 0;
      if (getLeaveListResult.links) {
        if (getLeaveListResult.links.prev) {
          prevUrl = parseInt(getLeaveListResult.links.prev, 10);
        }
        if (getLeaveListResult.links.next) {
          nextUrl = parseInt(getLeaveListResult.links.next, 10);
        }
      }

      let pageInfo = {
        'from': getLeaveListResult.meta.from || 0,
        'to': getLeaveListResult.meta.to || 0,
        'total': getLeaveListResult.meta.total ||0,
        'prev': prevUrl,
        'next': nextUrl,
      }
      setPagingInfo(pageInfo);
  }

  /**
   * 設定分頁
   */
  const setPagingInfo = (input) => {
    pagingInfo.from = input.from;
    pagingInfo.to = input.to;
    pagingInfo.total = input.total;
    pagingInfo.prev = input.prev;
    pagingInfo.next = input.next;
  }

  const result = ref("");
  /**
   * 同意請假申請
   */
  const approveLeave =  async (id) => {
    let leaveInfo = {
      "id": parseInt(id),
    }
    let applyLeavetResult = await api.leave.approveLeave(leaveInfo)
    if (applyLeavetResult.statusCode == 200) {
      getLeaveList();
    } else {
      result.value = `請假申請單號 ${id} 失敗`
    }
    
  };

  /**
   * 不同意請假申請
   */
  const disagreeLeave =  async (id) => {
    let applyLeavetResult = await api.leave.cancelLeave(id)
    if (applyLeavetResult.statusCode == 200) {
      getLeaveList();
    } else {
      result.value = `請假申請單號 ${id} 失敗`
    }
  };

  /**
   * 轉換時間格式
   * @param {*} input 
   */
  const formatTimeFormat = (input) => {
    let date = new Date(input);
    
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minute = String(date.getMinutes()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minute}`;
  }
</script>

<template>
  <section
    class="h-fit container mx-auto my-4 max-w-3xl p-8 border-2 rounded-md"
  >
    <div class="flex flex-col items-center">
      <div 
        class="w-4/6 h-40 m-auto mb-2 bg-white rounded-md shadow-lg border-2 border-slate-700 grid grid-cols-4 grid-rows-4 p-2 bg-gray-100"
        v-for="leave in leaveList"
      >
        <span class="col-span-2 row-span-1 row-start-1">
          <p class="text-left leading-4 text-sm">
            員工 {{ leave.name }}
          </p>
        </span>
        <span class="col-span-4 col-start-3 row-span-1 row-start-1 flex flex-col">
          <small class="text-right text-[12px]">
            開始時間：{{ formatTimeFormat(leave.startAt) }}
          </small>
          <small class="text-right text-[12px]">
            開始結束：{{ formatTimeFormat(leave.endAt) }}
          </small> 
        </span>
        <span class="col-span-1 row-span-1 row-start-4 relative">
          <small class="text-left text-[8px] float-left absolute bottom-0 left-0">
            申請單號: {{ leave.id }}
          </small> 
        </span>
        <span class="col-span-4 row-span-2 row-start-2">
          <p class="h-12 text-gray-700 text-base overflow-clip text-left">
            {{ leave.reason }}
          </p>
        </span>
        <span class="col-span-3 col-start-2 row-span-1 row-start-4 p-1 flex gap-4">
          <button
              class="h-full w-full  border-2 border-red-200 rounded-md bg-red-100"
              @click="disagreeLeave(leave.id)"
          >
            <img
              class="h-full w-full svg-red"
              src = "@/assets/icons/close.svg"
            />
          </button>
          <button
              class="h-full w-full border-2 border-green-200 rounded-md bg-green-100"
              @click="approveLeave(leave.id)"
            >
              <img
                class="h-full w-full svg-green"
                src = "@/assets/icons/check.svg"
              />
          </button>
        </span>
      </div>
    </div>
    <div class="flex flex-col items-center">
      <!-- Help text -->
      <span class="text-sm text-gray-700 dark:text-gray-400">
          目前為
          <span
            class="font-semibold text-gray-900 dark:text-white"
          >
            第{{ pagingInfo.from }} 件
          </span> 
          至 
          <span class="font-semibold text-gray-900 dark:text-white">
            第{{ pagingInfo.to }}  件
          </span>
            總數：
          <span class="font-semibold text-gray-900 dark:text-white">
            {{ pagingInfo.total }}
          </span>
      </span>
      <div class="inline-flex mt-2 xs:mt-0">
        <!-- Buttons -->
        <button 
          class="inline-flex items-center px-4 py-2 m-2 border-2 border-slate-700 text-sm font-medium"
          @click="getLeaveList(pagingInfo.prev)"
          v-if="pagingInfo.from > 1"
        >
          <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M7.707 14.707a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l2.293 2.293a1 1 0 010 1.414z" clip-rule="evenodd"></path></svg>
        </button>
        <button
          class="inline-flex items-center px-4 py-2 m-2 border-2 border-slate-700 text-sm font-medium"
          @click="getLeaveList(pagingInfo.next)"  
          v-if="pagingInfo.next !== 0"
        >
          <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
        </button>
      </div>
    </div>
  </section>
</template>

<style scoped>

</style>
