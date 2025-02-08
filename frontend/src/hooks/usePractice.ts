import { useReducer } from "react";
import { PracticeI, PracticeList } from "../model/practice";

type PracticeState = {
    [key: string]: PracticeI
}

enum PracticeActionKind {
    LIST = 'LIST',
    ONE = 'ONE',
  }

interface PracticeAction {
    type: PracticeActionKind
    itemList:  PracticeList
    item: PracticeI
} 


function practiceDispatch(state: PracticeState, action: PracticeAction): PracticeState{
    const {type, itemList, item} = action;
    switch (type) {
        case PracticeActionKind.LIST:
            for (let index = 0; index < itemList.practices.length; index++) {
                const element = itemList.practices[index];
                return {
                    ...state,
                    [element.practiceId]: element 
                }
            }
            return state
        case PracticeActionKind.ONE:
            return {
                ...state,
                [item.practiceId]: item
            }
        default:
            return state
    }

}

export function usePractices() {
    const [practices, updatePractices] = useReducer(practiceDispatch, null, ()=>{return {}});

    const getPractice = function(id: string){
        
    }
}