import 'package:flutter/material.dart';
import 'package:gosha_app/app/model_card.dart';
import 'package:waterfall_flow/waterfall_flow.dart';
import 'package:get/get.dart';
import 'package:gosha_app/controller/app_controller.dart';

class ModelList extends StatelessWidget {
  ModelList({Key? key}) : super(key: key);

  AppController controller = Get.find();

  @override
  Widget build(BuildContext context) {
    return Obx(() {
      return WaterfallFlow.builder(
        padding: EdgeInsets.all(5.0),
        itemCount: controller.models.length,
        gridDelegate: SliverWaterfallFlowDelegateWithMaxCrossAxisExtent(
          maxCrossAxisExtent: 480,
          crossAxisSpacing: 5.0,
          mainAxisSpacing: 5.0,
        ),
        itemBuilder: (BuildContext context, int index) {
          return ModelCard(index: index);
        },
      );
    });
  }
}
