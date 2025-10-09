export async function scenarioListLoader({ request }: { request: Request }): Promise<any> {
  const url = new URL(request.url);
  const queryName = url.searchParams.get('name');
  const rawTags = url.searchParams.get('tag');
  let tags: string[] | undefined = undefined;
  if (!!rawTags) {
    tags = rawTags.split(',');
  }
  return {}
//   const query: ScenarioQuery = { name: queryName || undefined, tag: tags }
//   try {
//     const response = await getAll(query);
//     return {
//       scenarios: response,
//       query
//     }
//   } catch (error: any) {
//     return {
//       scenarios: {
//         total: 0,
//         data: [],
//         pages: 0,
//         currentPage: 0
//       },
//       query
//     }
//   }
}

export async function scenarioItemLoader({ params }: { params: any }): Promise<any> {
  return {}
  // const response = await getScenario(params.scenarioId);
  // if (!response) {
  //   return {
  //     id: '',
  //     created_at: new Date().getTime(),
  //     name: '',
  //     payload: {
  //       // variables: [],
  //       blocks: [],
  //       edges: []
  //     },
  //     input_data: []
  //   }
  // }
  // return response;
}
