import * as cdk from "@aws-cdk/core";
import { MikrodStack } from "./src/stack";

const app = new cdk.App();

new MikrodStack(app, "MikrodStack");
