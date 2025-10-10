/**
 * Mock data for scenario flows (nodes and edges)
 * Based on existing mock data from /frontend/mock/scenario.mock.ts
 */

export const mockFlows: Record<string, { blocks: any[]; edges: any[] }> = {
  "1": {
    blocks: [
      { id: "O0geOBtqfndU4bHGSaLBK", type: "start", position: { x: 0, y: 0 } },
      {
        id: "CKbXtv4PCJxFNABlLM5Xr",
        type: "goto",
        position: { x: 255, y: -120 },
      },
      {
        id: "frCvQzBKsH6yUtEDLpNNo",
        type: "waitduration",
        position: { x: 255, y: 30 },
      },
      {
        id: "1NHlcJphR3cevDCG7dExe",
        type: "findelement",
        position: { x: 465, y: -60 },
      },
      {
        id: "vS07vrsyOjIzny9HkmgeD",
        type: "click",
        position: { x: 690, y: -75 },
      },
      {
        id: "fssj71wS0g_xwskFEjDHl",
        type: "findelement",
        position: { x: 900, y: -75 },
      },
      {
        id: "WRbaUo6OXN8EAlWtSjek_",
        type: "inputdata",
        position: { x: 1110, y: -75 },
      },
      {
        id: "HIG0NE21mvRkT9h6waB1W",
        type: "keypress",
        position: { x: 1095, y: 120 },
      },
      {
        id: "e3JZOUD6MZ0216GeZmnaO",
        type: "waitduration",
        position: { x: 1335, y: 0 },
      },
      {
        id: "w5UjxZ-_BrtT-xwou49LE",
        type: "screenshot",
        position: { x: 1560, y: 150 },
      },
      {
        id: "cWms0c3c8eLaTU8dSyQa-",
        type: "screenshot",
        position: { x: 690, y: 75 },
      },
      {
        id: "1Y_l5UoL0dSFETg3ouX8J",
        type: "click",
        position: { x: 1545, y: -45 },
      },
      {
        id: "S7gzi5NUB0cw8BInNrGyL",
        type: "screenshot",
        position: { x: 1815, y: 30 },
      },
      {
        id: "8_ccWPDjcdgxS97PAmYCB",
        type: "loaddata",
        position: { x: 15, y: 165 },
      },
    ],
    edges: [
      {
        id: "_AVbaRJJC0ZzcX68Tsfht",
        source: "O0geOBtqfndU4bHGSaLBK",
        target: "CKbXtv4PCJxFNABlLM5Xr",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "iK2lNRZEYRxnscvJSRRpB",
        source: "CKbXtv4PCJxFNABlLM5Xr",
        target: "frCvQzBKsH6yUtEDLpNNo",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "Q_t5UvE3qfGVPqBeDuhDa",
        source: "frCvQzBKsH6yUtEDLpNNo",
        target: "1NHlcJphR3cevDCG7dExe",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "z5CfPlXSNInzhzLfze2Z9",
        source: "1NHlcJphR3cevDCG7dExe",
        target: "vS07vrsyOjIzny9HkmgeD",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "Ni0g2tEOXmoxogWPJ_8E1",
        source: "vS07vrsyOjIzny9HkmgeD",
        target: "fssj71wS0g_xwskFEjDHl",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "8xMWjdWdQGkkvJ8SxFC0X",
        source: "fssj71wS0g_xwskFEjDHl",
        target: "WRbaUo6OXN8EAlWtSjek_",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "LkPLFsZNAqZXSBf5L_PfX",
        source: "WRbaUo6OXN8EAlWtSjek_",
        target: "HIG0NE21mvRkT9h6waB1W",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "KSJCgdR3_gxfiFTDUL44i",
        source: "HIG0NE21mvRkT9h6waB1W",
        target: "e3JZOUD6MZ0216GeZmnaO",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "-clkhTX_feraB9Ur8l3mB",
        source: "e3JZOUD6MZ0216GeZmnaO",
        target: "w5UjxZ-_BrtT-xwou49LE",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "kggJm_8IrDrCzrbt2WbP4",
        source: "1NHlcJphR3cevDCG7dExe",
        target: "cWms0c3c8eLaTU8dSyQa-",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "9p45OiyOyTD1ap6Og1vtd",
        source: "e3JZOUD6MZ0216GeZmnaO",
        target: "1Y_l5UoL0dSFETg3ouX8J",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "zvfk0VeBLZJKXdrJFmJNa",
        source: "1Y_l5UoL0dSFETg3ouX8J",
        target: "S7gzi5NUB0cw8BInNrGyL",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "ev7msD4ZivAVJol3Zwxsj",
        source: "8_ccWPDjcdgxS97PAmYCB",
        target: "CKbXtv4PCJxFNABlLM5Xr",
        sourceHandle: "node-target",
        targetHandle: "CKbXtv4PCJxFNABlLM5Xr-url",
      },
    ],
  },
  "2": {
    blocks: [
      { id: "start-1", type: "start", position: { x: 0, y: 0 } },
      { id: "goto-1", type: "goto", position: { x: 200, y: 0 } },
      { id: "screenshot-1", type: "screenshot", position: { x: 400, y: 0 } },
    ],
    edges: [
      {
        id: "edge-1",
        source: "start-1",
        target: "goto-1",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
      {
        id: "edge-2",
        source: "goto-1",
        target: "screenshot-1",
        sourceHandle: "node-target",
        targetHandle: "node-source",
      },
    ],
  },
  "3": {
    blocks: [
      { id: "start-3", type: "start", position: { x: 0, y: 0 } },
    ],
    edges: [],
  },
};

export const getFlowByScenarioId = (scenarioId: string) => {
  return mockFlows[scenarioId] || { blocks: [], edges: [] };
};
