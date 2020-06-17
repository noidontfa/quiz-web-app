/// <reference types="react-scripts" />
interface QuizInterface {
    id? : number;
    createdAt? : string;
    name? : string;
    description? : string;
    categoryRefer? : CategoryInterface;
    languageRefer? : LanguageInterface;
    timingRefer? : TimingInterface;
    stateRefer? : StateInterface;
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
    username? : string;
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

interface HistoryInterface {
    id? : number;
    numberRightAnswers? : number;
    score? : number;
    quizRefer? : QuizInterface;
    userRefer? : UserInterface;
    createAt? : string;
}

interface LanguageInterface {
    id? :number;
    name? : string;
}

interface StateInterface {
    id? :number;
    name? : string;
}

interface TokenInterface {
    token : string;
}