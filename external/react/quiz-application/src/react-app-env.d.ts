/// <reference types="react-scripts" />
interface QuizInterface {
    id? : number;
    createAt? : string;
    name? : string;
    description? : string;
    categoryRefer? : CategoryInterface;
    timingRefer? : TimingInterface;
    userRefer? : UserInterface;
    questionRefer? : Array<QuestionInterface>;
    totalQuestions? : number;
    ratings?: number;
    image? : string;

}
interface CategoryInterface {
    id? : number;
    name? : string;

}

interface TimingInterface {
    id? : number;
    name?: string;
    sec? : number;
}

interface UserInterface {
    id? : number;
    firstName?: string;
    lastName? : string;
    email? : string;
    dayOfBirth? : string;
    image? : string;
}

interface QuestionInterface {
    id? : number;
    name? : string;
    choices? : Array<ChoiceInterface>;
    image?: string;
}

interface ChoiceInterface {
    id? : number;
    name?: string;
    isRight?: boolean;
    image?: string;
}