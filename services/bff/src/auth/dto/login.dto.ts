import { Field, InputType, ObjectType } from '@nestjs/graphql'

@InputType()
export class LoginInput {
  @Field()
  email: string

  @Field()
  password: string
}

@ObjectType()
export class LoginOutput {
  @Field()
  message: string
}
