import { describe, test, expect, vi, beforeEach } from 'vitest'
import service from '../../api/index'
import httpObj from '../../request';


// function createFetchResponse(data) {
//     return { json: () => new Promise((resolve) => resolve(data)) }
// }
// vi.mock('axios', () => {
//     return {
//         default: {
//             get: vi.fn(),
//         },
//     }
// })

function createFetchResponse(data, m) {
        return 1
    }

describe.only('login test', () => {
    beforeEach(() => {
        //global.axios.mockReset()
    })

    test('makes a GET request to fetch todo list and returns the result', async () => {
        vi.mock('../../request', async () => {
            return {
                default: {
                    get: vi.fn().mockResolvedValue(createFetchResponse(1,3))
                }
            }
        })
        
        // httpObj.get.mockResolvedValue()
        //axios.get.mockResolvedValue(createFetchResponse(response))
        console.log(3)
        console.log(await service.login.getTreeList())
        // try {
        // const response =  await service.get('', '');
        // console.log(response)
        //     }
        //     catch (error) {
        //         console.log(error.response.statusText)
        //     //return Promise.reject(error);
        //     }
        
        //expect(todoList).toStrictEqual(response)
    })

})