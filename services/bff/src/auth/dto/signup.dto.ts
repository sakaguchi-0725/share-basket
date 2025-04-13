import { Field, InputType, ObjectType } from '@nestjs/graphql'

@InputType()
export class SignUpInput {
  @Field()
  name: string

  @Field()
  email: string

  @Field()
  password: string
}

@ObjectType()
export class SignUpOutput {
  @Field()
  message: string
}
