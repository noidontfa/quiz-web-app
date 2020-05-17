/// <reference types="react-scripts" />
interface QuizInterface {
    id? : number;
    createAt? : string;
    name? : string;
    description? : string;
    categoryRefer? : CategoryInterface;
    timingRefer? : TimingInterface;
    userRefer? : UserInterface;
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