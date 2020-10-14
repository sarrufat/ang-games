export interface ChessResult {
    done: boolean;
    ms: number;
    iterations: number;
    combinations: Combination[];
    numCombinations: number;
    msg: string;
    combLenght: number;
}

export interface Combination {
    positions: Position[];
}

export interface Position {
    piece: string;
    x: number;
    y: number;
}

export interface TaskId {
    taskId: string;
}