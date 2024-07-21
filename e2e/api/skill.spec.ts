import {test, expect} from '@playwright/test'

const apiUrlPrefix = 'http://localhost:8910/api/v1'

test.describe('Get skill by key', () => {
  test('should response one skill with status "success" when request GET /skills/:key', async ({
    request,
  }) => {
    await request.post(apiUrlPrefix + '/skills',
      {
        data: {
          key: 'python',
          name: 'Python',
          description: 'Python is an interpreted, high-level, general-purpose programming language.',
          logo: 'https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg',
          tags: ['programming language', 'scripting']
        }
      }
    )

    const resp = await request.get(apiUrlPrefix + '/skills/python')
  
    expect(resp.ok()).toBeTruthy()
    expect(await resp.json()).toEqual(
      expect.objectContaining({
        status: 'success',
        data: {
          key: 'python',
          name: 'Python',
          description: expect.any(String),
          logo: expect.any(String),
          tags: expect.arrayContaining(['programming language', 'scripting']),
        },
      })
    )

    await request.delete(apiUrlPrefix + '/skills/python')
  })
})